import { Badge } from '@/components/ui/badge';
import { Accident } from '../types/aviationTypes';

interface AccidentBadgeProps {
  accident: Accident;
}

// Component to render badges based on injury data and specific remark text conditions
const AccidentBadges: React.FC<AccidentBadgeProps> = ({ accident }) => {
  const { injuries, remark_text } = accident;

  // Initialize injury counts for each severity and person type
  const injuryCounts: { [key: string]: number } = {
    ground_fatal: 0,
    ground_serious: 0,
    ground_minor: 0,
    ground_unknown: 0,
    flight_crew_fatal: 0,
    flight_crew_serious: 0,
    flight_crew_minor: 0,
    flight_crew_unknown: 0,
    cabin_crew_fatal: 0,
    cabin_crew_serious: 0,
    cabin_crew_minor: 0,
    cabin_crew_unknown: 0,
    passengers_fatal: 0,
    passengers_serious: 0,
    passengers_minor: 0,
    passengers_unknown: 0,
  };

  // Iterate through injuries and accumulate counts
  injuries?.forEach((injury) => {
    const key = `${injury.person_type}_${injury.injury_severity}`;
    injuryCounts[key] += injury.count;
  });

  // Function to determine label based on count and pluralization
  const getLabel = (count: number, singular: string, plural: string) =>
    count === 1 ? singular : plural;

  return (
    <>
      {/* Render badges for injuries */}
      {injuryCounts.ground_fatal +
        injuryCounts.flight_crew_fatal +
        injuryCounts.cabin_crew_fatal +
        injuryCounts.passengers_fatal >
        0 && (
        <Badge className={`bg-red-500 mb-1`}>{`${
          injuryCounts.ground_fatal +
          injuryCounts.flight_crew_fatal +
          injuryCounts.cabin_crew_fatal +
          injuryCounts.passengers_fatal
        } ${getLabel(
          injuryCounts.ground_fatal +
            injuryCounts.flight_crew_fatal +
            injuryCounts.cabin_crew_fatal +
            injuryCounts.passengers_fatal,
          'Fatality',
          'Fatalities'
        )}`}</Badge>
      )}
      {injuryCounts.ground_serious +
        injuryCounts.flight_crew_serious +
        injuryCounts.cabin_crew_serious +
        injuryCounts.passengers_serious >
        0 && (
        <Badge className={`bg-yellow-500 mb-1`}>{`${
          injuryCounts.ground_serious +
          injuryCounts.flight_crew_serious +
          injuryCounts.cabin_crew_serious +
          injuryCounts.passengers_serious
        } ${getLabel(
          injuryCounts.ground_serious +
            injuryCounts.flight_crew_serious +
            injuryCounts.cabin_crew_serious +
            injuryCounts.passengers_serious,
          'Serious Injury',
          'Serious Injuries'
        )}`}</Badge>
      )}
      {injuryCounts.ground_minor +
        injuryCounts.flight_crew_minor +
        injuryCounts.cabin_crew_minor +
        injuryCounts.passengers_minor >
        0 && (
        <Badge className={`bg-green-500 mb-1`}>{`${
          injuryCounts.ground_minor +
          injuryCounts.flight_crew_minor +
          injuryCounts.cabin_crew_minor +
          injuryCounts.passengers_minor
        } ${getLabel(
          injuryCounts.ground_minor +
            injuryCounts.flight_crew_minor +
            injuryCounts.cabin_crew_minor +
            injuryCounts.passengers_minor,
          'Minor Injury',
          'Minor Injuries'
        )}`}</Badge>
      )}

      {/* Additional badges based on specific conditions found in remark text */}
      {(remark_text.toLowerCase().includes('bird') ||
        remark_text.toLowerCase().includes('birds')) && (
        <Badge className="bg-blue-500 mb-1">Birds</Badge>
      )}
      {(remark_text.toLowerCase().includes('stall') ||
        remark_text.toLowerCase().includes('stalled')) && (
        <Badge className="bg-purple-500 mb-1">Stall</Badge>
      )}
      {remark_text.toLowerCase().includes('fire') && (
        <Badge className="bg-orange-500 mb-1">Fire</Badge>
      )}
    </>
  );
};

export default AccidentBadges;
