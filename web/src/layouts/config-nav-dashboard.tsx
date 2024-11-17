import { paths } from 'src/routes/paths';

import { CONFIG } from 'src/config-global';

import { SvgColor } from 'src/components/svg-color';

// ----------------------------------------------------------------------

const icon = (name: string) => {
  const src = `${CONFIG.assetsDir}/assets/icons/navbar/${name}.svg`;
  return <SvgColor src={src} />;
};

export const ICONS = {
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
  history: icon('history'),
};

export const ICONS_SRC = {
  job: `${CONFIG.assetsDir}/assets/icons/navbar/ic-job.svg`,
  blog: `${CONFIG.assetsDir}/assets/icons/navbar/ic-blog.svg`,
  chat: `${CONFIG.assetsDir}/assets/icons/navbar/ic-chat.svg`,
  mail: `${CONFIG.assetsDir}/assets/icons/navbar/ic-mail.svg`,
  user: `${CONFIG.assetsDir}/assets/icons/navbar/ic-user.svg`,
  file: `${CONFIG.assetsDir}/assets/icons/navbar/ic-file.svg`,
  lock: `${CONFIG.assetsDir}/assets/icons/navbar/ic-lock.svg`,
  tour: `${CONFIG.assetsDir}/assets/icons/navbar/ic-tour.svg`,
  order: `${CONFIG.assetsDir}/assets/icons/navbar/ic-order.svg`,
  label: `${CONFIG.assetsDir}/assets/icons/navbar/ic-label.svg`,
  blank: `${CONFIG.assetsDir}/assets/icons/navbar/ic-blank.svg`,
  kanban: `${CONFIG.assetsDir}/assets/icons/navbar/ic-kanban.svg`,
  folder: `${CONFIG.assetsDir}/assets/icons/navbar/ic-folder.svg`,
  course: `${CONFIG.assetsDir}/assets/icons/navbar/ic-course.svg`,
  banking: `${CONFIG.assetsDir}/assets/icons/navbar/ic-banking.svg`,
  booking: `${CONFIG.assetsDir}/assets/icons/navbar/ic-booking.svg`,
  invoice: `${CONFIG.assetsDir}/assets/icons/navbar/ic-invoice.svg`,
  product: `${CONFIG.assetsDir}/assets/icons/navbar/ic-product.svg`,
  calendar: `${CONFIG.assetsDir}/assets/icons/navbar/ic-calendar.svg`,
  disabled: `${CONFIG.assetsDir}/assets/icons/navbar/ic-disabled.svg`,
  external: `${CONFIG.assetsDir}/assets/icons/navbar/ic-external.svg`,
  menuItem: `${CONFIG.assetsDir}/assets/icons/navbar/ic-menu-item.svg`,
  ecommerce: `${CONFIG.assetsDir}/assets/icons/navbar/ic-ecommerce.svg`,
  analytics: `${CONFIG.assetsDir}/assets/icons/navbar/ic-analytics.svg`,
  dashboard: `${CONFIG.assetsDir}/assets/icons/navbar/ic-dashboard.svg`,
  parameter: `${CONFIG.assetsDir}/assets/icons/navbar/ic-parameter.svg`,
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
      { title: 'Maas', path: paths.dashboard.maas, icon: ICONS.ecommerce },
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
