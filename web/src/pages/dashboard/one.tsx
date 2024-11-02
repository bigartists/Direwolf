import { Helmet } from 'react-helmet-async';

import { CONFIG } from 'src/config-global';

import { BlankView } from 'src/sections/blank/view';

import { MultiLLMChat } from 'src/sections/chat/view';

// ----------------------------------------------------------------------

const metadata = { title: `Page one | Dashboard - ${CONFIG.appName}` };

export default function Page() {
  return (
    <>
      <Helmet>
        <title> {metadata.title}</title>
      </Helmet>
      <MultiLLMChat title="MultiLLMChat page" />
    </>
  );
}
