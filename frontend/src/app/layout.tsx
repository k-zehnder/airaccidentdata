import './globals.css';
import { ThemeProvider } from '@/components/theme-provider';
import { cn } from '../lib/utils';

const fontStyling = {
  fontFamily: '"Inter", sans-serif',
};

interface Metadata {
  title: string;
  description: string;
  faviconUrl: string;
  ogImageUrl: string;
}

export const metadata: Metadata = {
  title: 'airaccidentdata.com',
  description: 'Latest plane crash accidents and incident news!',
  faviconUrl: 'https://s.airaccidentdata.com/favicon.ico',
  ogImageUrl: 'https://s.airaccidentdata.com/og-image.png',
};

interface RootLayoutProps {
  children: React.ReactNode;
}

export default async function RootLayout({ children }: RootLayoutProps) {
  return (
    <html lang="en">
      <head>
        <title>{metadata.title}</title>
        <meta name="description" content={metadata.description} />
        <meta property="og:title" content={metadata.title} />
        <meta property="og:description" content={metadata.description} />
        <meta property="og:image" content={metadata.ogImageUrl} />
        <link
          rel="icon"
          href={metadata.faviconUrl}
          type="image/x-icon"
          sizes="16x16"
        />
        <script
          async
          src="https://www.googletagmanager.com/gtag/js?id=G-D9DT897WDG"
        ></script>
        <script
          dangerouslySetInnerHTML={{
            __html: `
              window.dataLayer = window.dataLayer || [];
              function gtag(){dataLayer.push(arguments);}
              gtag('js', new Date());
              gtag('config', 'G-D9DT897WDG', {
                page_path: window.location.pathname,
              });
            `,
          }}
        />
      </head>
      <body
        className={cn('min-h-screen bg-background antialiased')}
        style={fontStyling}
      >
        <ThemeProvider attribute="class" defaultTheme="system" enableSystem>
          <div className="relative flex min-h-screen flex-col">
            <div className="flex w-full flex-1 flex-col overflow-hidden">
              {children}
            </div>
          </div>
        </ThemeProvider>
      </body>
    </html>
  );
}
