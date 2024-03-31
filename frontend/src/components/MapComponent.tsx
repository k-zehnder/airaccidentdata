import React, { useRef, useEffect } from 'react';
import * as d3 from 'd3';
import { Topology, GeometryCollection } from 'topojson-specification';
import { feature } from 'topojson-client';

const MapComponent: React.FC = () => {
  const svgRef = useRef<SVGSVGElement>(null);

  useEffect(() => {
    // Define the type for your specific TopoJSON structure
    type USATopoJson = Topology<{ states: GeometryCollection }>;

    const width = 975;
    const height = 610;
    const projection = d3.geoAlbersUsa().scale(1300).translate([width / 2, height / 2]);
    const path = d3.geoPath(projection);

    d3.json<USATopoJson>('/states-10m.json').then((us) => {
      if (us && svgRef.current) {
        const states = feature(us, us.objects.states as GeometryCollection).features;

        const svg = d3.select(svgRef.current)
          .attr('viewBox', [0, 0, width, height]);

        svg.selectAll('path')
          .data(states)
          .enter().append('path')
            .attr('fill', '#444')
            .attr('d', path);

        // Hardcoded coordinates for New York, NY
        const nyCoords: [number, number] = [-74.0060, 40.7128];
        const projectedNYCoords = projection(nyCoords);

        if (projectedNYCoords) {
          svg.append('circle')
            .attr('cx', projectedNYCoords[0])
            .attr('cy', projectedNYCoords[1])
            .attr('r', 5)
            .attr('fill', 'red');
        }
      }
    });
  }, []);

  return (
    <svg ref={svgRef}></svg>
  );
};

export default MapComponent;
