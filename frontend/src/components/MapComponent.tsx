import React, { useRef, useEffect } from 'react';
import * as d3 from 'd3';
import { Topology, GeometryCollection } from 'topojson-specification';
import { feature } from 'topojson-client';

// Props interface to define expected props for the component
interface MapComponentProps {
  latitude?: number; // Optional latitude for zooming and centering on a specific point
  longitude?: number; // Optional longitude for zooming and centering on a specific point
}

const MapComponent: React.FC<MapComponentProps> = ({ latitude, longitude }) => {
  // useRef to keep a reference to the SVG element in the DOM
  const svgRef = useRef<SVGSVGElement>(null);

  useEffect(() => {
    // Type definition for the topojson data structure
    type USATopoJson = Topology<{ states: GeometryCollection }>;

    // Dimensions of the SVG canvas
    const width = 975;
    const height = 610;

    // Setting up the geographic projection for the USA map
    let projection = d3
      .geoAlbersUsa()
      .scale(1300)
      .translate([width / 2, height / 2]);
    let path = d3.geoPath(projection);

    // Function to update projection based on provided coordinates, enabling zoom
    const updateProjection = (coords: [number, number]) => {
      const zoomLevel = 4000; // High zoom level for focusing closely on a point
      projection = d3
        .geoAlbersUsa()
        .scale(zoomLevel)
        .translate([width / 2, height / 2]);

      const point = projection(coords);
      if (point) {
        // Check if the projected point is within the map
        projection.translate([
          width / 2 - point[0] + width / 2,
          height / 2 - point[1] + height / 2,
        ]);
        path = d3.geoPath().projection(projection);
      }
    };

    // Load geographic data from a topojson file
    d3.json<USATopoJson>('/states-10m.json').then((us) => {
      if (us && svgRef.current) {
        // Extract the state features from the topojson
        const states = feature(
          us,
          us.objects.states as GeometryCollection
        ).features;

        // If latitude and longitude are provided, update the map's center and zoom
        if (latitude !== undefined && longitude !== undefined) {
          updateProjection([longitude, latitude]);
        }

        // Select the SVG element via d3 and clear any previous content
        const svg = d3.select(svgRef.current).html('');
        svg
          .attr('viewBox', [0, 0, width, height]) // Set the viewBox for responsive scaling
          .selectAll('path') // Select all path elements for drawing states
          .data(states) // Bind state data to path elements
          .enter()
          .append('path') // Append a path element for each state
          .attr('fill', '#444') // Set the fill color for each state
          .attr('d', path) // Generate path data using geoPath
          .attr('stroke', '#fff'); // Set the stroke color for state boundaries

        // Add text labels for each state at their geographic centroid
        svg
          .selectAll('text')
          .data(states)
          .enter()
          .append('text')
          // @ts-ignore
          .filter((d) => {
            const centroid = projection(d3.geoCentroid(d));
            return (
              centroid &&
              centroid[0] >= 0 &&
              centroid[0] <= width &&
              centroid[1] >= 0 &&
              centroid[1] <= height
            );
          })
          .attr('x', (d) => {
            const centroid = projection(d3.geoCentroid(d)); // Calculate centroid for label positioning
            return centroid ? centroid[0] : null; // Use the x-coordinate of the centroid
          })
          .attr('y', (d) => {
            const centroid = projection(d3.geoCentroid(d)); // Calculate centroid for label positioning
            return centroid ? centroid[1] : null; // Use the y-coordinate of the centroid
          })
          .attr('text-anchor', 'middle') // Center the text horizontally
          .attr('alignment-baseline', 'central') // Center the text vertically
          .attr('fill', 'white') // Set text color
          .style('font-size', '20px') // Set text size
          .text((d: any) => d.properties.name); // Set the text to the state's name

        // Add a red dot at the specified latitude and longitude
        if (latitude !== undefined && longitude !== undefined) {
          const coords: [number, number] = [longitude, latitude];
          const projectedCoords = projection(coords);
          if (projectedCoords) {
            // Check if coordinates are within the visible map
            svg
              .append('circle') // Add a circle to mark the location
              .attr('cx', projectedCoords[0])
              .attr('cy', projectedCoords[1])
              .attr('r', 15) // Radius of the circle
              .attr('fill', 'red'); // Color of the circle
          }
        }
      }
    });
  }, [latitude, longitude]); // Depend on latitude and longitude to re-run effects when they change

  return <svg ref={svgRef}></svg>; // Render the SVG element
};

export default MapComponent;
