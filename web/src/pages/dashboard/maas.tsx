import { Helmet } from 'react-helmet-async';

import { CONFIG } from 'src/config-global';

import { ModelView } from 'src/sections/maas/view';

// ----------------------------------------------------------------------

const metadata = { title: `Page Maas | Dashboard - ${CONFIG.appName}` };

export default function Page() {
  return (
    <>
      <Helmet>
        <title> {metadata.title}</title>
      </Helmet>

      <ModelView title="Maas列表" />
    </>
  );
}
