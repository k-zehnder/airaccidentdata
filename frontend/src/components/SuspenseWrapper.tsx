import React, { Suspense, ReactNode } from 'react';
import Loader from '@/components/Loader';

interface SuspenseWrapperProps {
  children: ReactNode;
}

// Component to wrap children with Suspense, showing Loader while loading
const SuspenseWrapper: React.FC<SuspenseWrapperProps> = ({ children }) => (
  <Suspense fallback={<Loader />}>{children}</Suspense>
);

export default SuspenseWrapper;
