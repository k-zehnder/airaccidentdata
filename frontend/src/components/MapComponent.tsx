import React, { useRef, useEffect, useState } from 'react';
import maplibregl from 'maplibre-gl';
import { withIdentityPoolId } from "@aws/amazon-location-utilities-auth-helper";

interface MapComponentProps {
  latitude: number;
  longitude: number;
}

const MapComponent: React.FC<MapComponentProps> = ({ latitude, longitude }) => {
  const [mapInitialized, setMapInitialized] = useState<boolean>(false);
  const mapContainer = useRef<HTMLDivElement>(null);
  const mapInstance = useRef<maplibregl.Map | null>(null);
  const circleLayerId = 'circle-layer';

  useEffect(() => {
    async function initializeMap() {
      const identityPoolId = process.env.NEXT_PUBLIC_IDENTITY_POOL_ID || '';
      const mapName = process.env.NEXT_PUBLIC_MAP_NAME || '';
      const region = process.env.NEXT_PUBLIC_REGION || '';

      const authHelper = await withIdentityPoolId(identityPoolId);

      if (!mapContainer.current || mapInitialized) return;

      const map = new maplibregl.Map({
        container: mapContainer.current,
        center: [longitude, latitude], // Original coordinates
        zoom: 10,
        style: `https://maps.geo.${region}.amazonaws.com/maps/v0/maps/${mapName}/style-descriptor`,
        ...authHelper.getMapAuthenticationOptions(),
      });

      map.addControl(new maplibregl.NavigationControl(), 'top-left');

      mapInstance.current = map;
      setMapInitialized(true);

      map.on('load', () => {
        // Add a circle layer at the center of the map
        map.addLayer({
          'id': circleLayerId,
          'type': 'circle',
          'source': {
            'type': 'geojson',
            'data': {
              'type': 'Feature',
              'geometry': {
                'type': 'Point',
                'coordinates': [longitude, latitude]
              }
            } as GeoJSON.Feature<GeoJSON.Point> 
          },
          'paint': {
            'circle-radius': 20, 
            'circle-color': '#FF0000' 
          }
        });
      });
    }

    initializeMap();

    return () => {
      if (mapInstance.current) {
        mapInstance.current.remove();
        mapInstance.current = null;
        setMapInitialized(false);
      }
    };
  }, [mapInitialized]);

  return <div ref={mapContainer} style={{ width: '100%', height: '500px', overflow: 'hidden' }} />;
};

export default MapComponent;
