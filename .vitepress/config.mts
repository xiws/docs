import { defineConfig } from 'vitepress'
import {sidebarConfig} from './sidebarConfig'

var list='[{text:database,items:[{text: test,link:./docs/database/test.md}]},{text:test,items:[{text:test.md,link:./docs/test/test.md}]}]';

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "docs",
  description: "A VitePress Site",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    logo: 'assets/logo.svg',
    search: {
      provider: 'local'
    },
    nav: [
      { text: 'Home', link: '/' },
      {
        text: 'Dropdown Menu',
        items: [
          { text: 'Item A', link: '/item-1' },
          { text: 'Item B', link: '/item-2' },
          { text: 'Item C', link: '/item-3' }
        ]
      }
    ],

    sidebar: sidebarConfig,

    socialLinks: [
      { icon: 'github', link: 'https://github.com/xiws/docs' }
    ]
  },
  outDir: "./docs",
  base: '/docs', // 设置 base 为相对路径
})
