// Register `hName`, `hProperties` types, used when turning markdown to HTML:
/// <reference types="mdast-util-to-hast" />
// Register directive nodes in mdast:
/// <reference types="mdast-util-directive" />

import { visit } from 'unist-util-visit';

// This plugin is an example to turn `::youtube` into iframes.
function myRemarkPlugin() {
  /**
   * @param {import('mdast').Root} tree
   *   Tree.
   * @param {import('vfile').VFile} file
   *   File.
   * @returns {undefined}
   *   Nothing.
   */
  return (tree: any, file: any) => {
    visit(tree, (node) => {
      if (
        node.type === 'containerDirective' ||
        node.type === 'leafDirective' ||
        node.type === 'textDirective'
      ) {
        if (node.name !== 'youtube') return;

        const data = node.data || (node.data = {});
        const attributes = node.attributes || {};
        const id = attributes.id;

        if (node.type === 'textDirective') {
          file.fail(
            'Unexpected `:youtube` text directive, use two colons for a leaf directive',
            node
          );
        }

        if (!id) {
          file.fail('Unexpected missing `id` on `youtube` directive', node);
        }

        data.hName = 'iframe';
        data.hProperties = {
          src: `https://www.youtube.com/embed/${id}`,
          width: 200,
          height: 200,
          frameBorder: 0,
          allow: 'picture-in-picture',
          allowFullScreen: true,
        };
      }
    });
  };
}
