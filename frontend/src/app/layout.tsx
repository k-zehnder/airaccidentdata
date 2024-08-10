import './globals.css';
import { ThemeProvider } from '@/components/theme-provider';
import { cn } from '../lib/utils';
import { GoogleAnalytics } from '@next/third-parties/google';
import { AccidentProvider } from '@/contexts/AccidentContext';

const fontStyling = {
  fontFamily: '"Inter", sans-serif',
};

interface RootLayoutProps {
  children: React.ReactNode;
}

const RootLayout: React.FC<RootLayoutProps> = ({ children }) => {
  return (
    <html lang="en">
      <head>
        <title>airaccidentdata.com</title>
        <meta
          name="description"
          content="Latest plane crash accidents and incident news!"
        />
        <meta property="og:title" content="airaccidentdata.com" />
        <meta
          property="og:description"
          content="Latest plane crash accidents and incident news!"
        />
        <meta
          property="og:image"
          content="https://s.airaccidentdata.com/og-image.png"
        />
        <link
          rel="icon"
          href="https://s.airaccidentdata.com/favicon.ico"
          type="image/x-icon"
          sizes="16x16"
        />
      </head>
      <body
        className={cn('min-h-screen bg-background antialiased')}
        style={fontStyling}
      >
        <ThemeProvider attribute="class" defaultTheme="system" enableSystem>
          <AccidentProvider>
            {' '}
            {/* Wrap children with AccidentProvider */}
            <div className="relative flex min-h-screen flex-col">
              <div className="flex w-full flex-1 flex-col overflow-hidden">
                {children}
              </div>
            </div>
          </AccidentProvider>
        </ThemeProvider>
        <GoogleAnalytics gaId="G-D9DT897WDG" />
      </body>
    </html>
  );
};

export default RootLayout;
