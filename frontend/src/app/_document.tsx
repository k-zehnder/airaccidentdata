import Document, { Html, Head, Main, NextScript } from 'next/document';
import { metadata } from './layout';

class MyDocument extends Document {
  render() {
    return (
      <Html lang="en">
        <Head>
          {/* Define meta tags */}
          <meta property="og:title" content={metadata.title} />
          <meta property="og:image" content={metadata.imageUrl} />
        </Head>
        <body>
          <Main />
          <NextScript />
        </body>
      </Html>
    );
  }
}

export default MyDocument;
