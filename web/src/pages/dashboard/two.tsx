import { Helmet } from 'react-helmet-async';

import { CONFIG } from 'src/config-global';

import { ModelView } from 'src/sections/model/view';

// ----------------------------------------------------------------------

const metadata = { title: `Page two | Dashboard - ${CONFIG.appName}` };

export default function Page() {
  return (
    <>
      <Helmet>
        <title> {metadata.title}</title>
      </Helmet>

      <ModelView title="模型列表" />
    </>
  );
}
