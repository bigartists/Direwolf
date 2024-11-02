import { paths } from 'src/routes/paths';

import { CONFIG } from 'src/config-global';

import { SvgColor } from 'src/components/svg-color';

// ----------------------------------------------------------------------

const icon = (name: string) => {
  const src = `${CONFIG.assetsDir}/assets/icons/navbar/${name}.svg`;
  console.log('🚀 ~ icon ~ src:', src);

  return <SvgColor src={src} />;
};

const ICONS = {
  job: icon('ic-job'),
  blog: icon('ic-blog'),
  chat: icon('ic-chat'),
  mail: icon('ic-mail'),
  user: icon('ic-user'),
  file: icon('ic-file'),
  lock: icon('ic-lock'),
  tour: icon('ic-tour'),
  order: icon('ic-order'),
  label: icon('ic-label'),
  blank: icon('ic-blank'),
  kanban: icon('ic-kanban'),
  folder: icon('ic-folder'),
  course: icon('ic-course'),
  banking: icon('ic-banking'),
  booking: icon('ic-booking'),
  invoice: icon('ic-invoice'),
  product: icon('ic-product'),
  calendar: icon('ic-calendar'),
  disabled: icon('ic-disabled'),
  external: icon('ic-external'),
  menuItem: icon('ic-menu-item'),
  ecommerce: icon('ic-ecommerce'),
  analytics: icon('ic-analytics'),
  dashboard: icon('ic-dashboard'),
  parameter: icon('ic-parameter'),
};

// ----------------------------------------------------------------------

export const navData = [
  /**
   * Overview
   */
  {
    subheader: 'basic ability',
    items: [
      { title: 'Chats', path: paths.dashboard.root, icon: ICONS.dashboard },
      { title: 'Models', path: paths.dashboard.models, icon: ICONS.ecommerce },
      // { title: 'Batch inference', path: paths.dashboard.three, icon: ICONS.analytics },
    ],
  },
  /**
   * Management
   */
  // {
  //   subheader: 'RAG',
  //   items: [
  //     { title: 'vector database', path: paths.dashboard.group.root, icon: ICONS.user },
  //     { title: 'embedding Models', path: paths.dashboard.group.five, icon: ICONS.user },
  //     { title: 'segmentation', path: paths.dashboard.group.five, icon: ICONS.user },
  //     { title: 'recall', path: paths.dashboard.group.six, icon: ICONS.user },
  //   ],
  // },
  // {
  //   subheader: 'Agent',
  //   items: [
  //     { title: 'Debugger', path: paths.dashboard.group.root, icon: ICONS.user },
  //     { title: 'App', path: paths.dashboard.group.five, icon: ICONS.user },
  //     { title: 'Share', path: paths.dashboard.group.six, icon: ICONS.user },
  //   ],
  // },
  // {
  //   subheader: 'Agent',
  //   items: [
  //     {
  //       title: 'Group',
  //       path: paths.dashboard.group.root,
  //       icon: ICONS.user,
  //       children: [
  //         { title: 'Four', path: paths.dashboard.group.root },
  //         { title: 'Five', path: paths.dashboard.group.five },
  //         { title: 'Six', path: paths.dashboard.group.six },
  //       ],
  //     },
  //   ],
  // },
];
