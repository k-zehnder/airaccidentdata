import React, { useRef, useEffect } from 'react';
import maplibregl from 'maplibre-gl';
import dotenv from 'dotenv';

dotenv.config();

interface MapComponentProps {
  latitude: number;
  longitude: number;
}

const MapComponent: React.FC<MapComponentProps> = ({ latitude, longitude }) => {
  const mapContainer = useRef<HTMLDivElement>(null);
  const mapInstance = useRef<maplibregl.Map | null>(null); 
  const circleLayerId = 'circle-layer';

  useEffect(() => {
    const apiKey = process.env.NEXT_PUBLIC_AWS_LOCATION_API_KEY; 

    if (mapContainer.current && apiKey && !mapInstance.current) {
      mapInstance.current = new maplibregl.Map({
        container: mapContainer.current,
        center: [longitude, latitude],
        zoom: 16,
        style: `https://maps.geo.us-west-2.amazonaws.com/maps/v0/maps/airaccidentdatamap/style-descriptor?key=${apiKey}`
      });

      // Add navigation control to the top-left corner of the map
      mapInstance.current.addControl(new maplibregl.NavigationControl(), 'top-left');

      mapInstance.current.on('load', () => {
        // Add a circle layer at the center of the map
        mapInstance.current?.addLayer({
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
            'circle-radius': 50, // in meters
            'circle-color': '#FF0000' // red color
          }
        });
      });
    }

    // Cleanup function to remove the map instance and circle layer when component unmounts
    return () => {
      if (mapInstance.current) {
        mapInstance.current.remove();
        mapInstance.current = null;
      }
    };
  }, [latitude, longitude]);

  return <div ref={mapContainer} style={{ width: '100%', height: '500px', overflow: 'hidden' }} />;
};

export default MapComponent;
