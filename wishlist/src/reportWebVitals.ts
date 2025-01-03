import { onLCP, onINP, onCLS, Metric } from 'web-vitals';

const reportWebVitals = (onPerfEntry?: (metric: Metric) => void) => {
  if (onPerfEntry && typeof onPerfEntry === 'function') {
    onCLS(onPerfEntry);
    onINP(onPerfEntry);
    onCLS(onPerfEntry);
  }
};

export default reportWebVitals;
