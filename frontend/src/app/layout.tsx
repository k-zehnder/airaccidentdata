import "./globals.css";
import { ThemeProvider } from "@/components/theme-provider";
import { cn } from "../lib/utils";
import Head from 'next/head'; 

// Define the font styling using CSS-in-JS or a separate CSS class
const fontStyling = {
  fontFamily: '"Inter", sans-serif',
};

export const metadata = {
  title: "airaccidentdata.com",
  description: "Latest plane crash accidents and incident news!",
  faviconUrl: "http://s.airaccidentdata.com/og-image.png",
};

interface RootLayoutProps {
  children: React.ReactNode;
}

export default async function RootLayout({ children }: RootLayoutProps) {

  return (
      <html lang="en">
        <body
          className={cn("min-h-screen bg-background antialiased")}
          style={fontStyling}
        >
          <ThemeProvider attribute="class" defaultTheme="system" enableSystem>
            <div className="relative flex min-h-screen flex-col">
              <div className="flex w-full flex-1 flex-col overflow-hidden">
                <link rel="icon" href="https://s.airaccidentdata.com/favicon.ico" type="image/xicon" sizes="16x16" />
                {children}
              </div>
            </div>
          </ThemeProvider>
        </body>
      </html>
   );
}
