/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],

  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],

  theme: {
    fontFamily: {
      // https://github.com/system-fonts/modern-font-stacks
      humanist: [
        'Seravek',
        'Gill Sans Nova',
        'Ubuntu',
        'Calibri',
        'DejaVu Sans',
        'source-sans-pro',
        'sans-serif',
      ],
      monocode: [
        'ui-monospace',
        'Cascadia Code',
        'Source Code Pro',
        'Menlo',
        'Consolas',
        'DejaVu Sans Mono',
        'monospace',
      ],
    },
    extend: {},
  },

  plugins: [require('daisyui')],

  daisyui: {
    styled: true,

    themes: [
      {
        mytheme: {
          primary: '#bfdbfe',
          secondary: '#463aa1',
          accent: '#c149ad',
          neutral: '#021431',
          'base-100': '#ffffff',
          info: '#93e6fb',
          success: '#80ced1',
          warning: '#efd8bd',
          error: '#e58b8b',
        },
      },
      'luxury',
    ],
    darkTheme: 'luxury',
    base: true,
    utils: true,
    logs: true,
    rtl: false,
    prefix: '',
  },
}
