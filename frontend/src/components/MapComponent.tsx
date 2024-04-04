import React, { useRef, useEffect } from 'react';
import * as d3 from 'd3';
import { Topology, GeometryCollection } from 'topojson-specification';
import { feature } from 'topojson-client';

interface MapComponentProps {
  latitude?: number;
  longitude?: number;
}

const MapComponent: React.FC<MapComponentProps> = ({ latitude, longitude }) => {
  const svgRef = useRef<SVGSVGElement>(null);

  useEffect(() => {
    type USATopoJson = Topology<{ states: GeometryCollection }>;

    const width = 975;
    const height = 610;
    let projection = d3.geoAlbersUsa().scale(1300).translate([width / 2, height / 2]);
    let path = d3.geoPath(projection);

    const updateProjection = (coords: [number, number]) => {
      const zoomLevel = 4000; // Adjust this value to control the zoom level
      projection = d3.geoAlbersUsa().scale(zoomLevel).translate([width / 2, height / 2]);

      // Calculate the point on the map to center on
      const point = projection(coords);
      if (point) {
        // Adjust the projection's translation
        projection.translate([
          width / 2 - point[0] + width / 2,
          height / 2 - point[1] + height / 2
        ]);
        path = d3.geoPath().projection(projection);
      }
    };

    d3.json<USATopoJson>('/states-10m.json').then((us) => {
      if (us && svgRef.current) {
        const states = feature(us, us.objects.states as GeometryCollection).features;

        // Clear the previous contents
        const svg = d3.select(svgRef.current).html('');

        // Check if latitude and longitude are provided to update the projection
        if (latitude !== undefined && longitude !== undefined) {
          updateProjection([longitude, latitude]);
        }

        // Draw the states
        svg.attr('viewBox', [0, 0, width, height])
           .selectAll('path')
           .data(states)
           .enter().append('path')
           .attr('fill', '#444')
           .attr('d', path);

        // Draw the red dot for the accident location
        if (latitude !== undefined && longitude !== undefined) {
          const coords: [number, number] = [longitude, latitude];
          const projectedCoords = projection(coords);
          if (projectedCoords) {
            svg.append('circle')
               .attr('cx', projectedCoords[0])
               .attr('cy', projectedCoords[1])
               .attr('r', 5)
               .attr('fill', 'red');
          }
        }
      }
    });
  }, [latitude, longitude]); // Re-run when latitude or longitude changes

  return <svg ref={svgRef}></svg>;
};

export default MapComponent;
