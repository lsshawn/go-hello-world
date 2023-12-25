/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./view/**/*.{html,js,templ,go}",
    "./view/common/**/*.{html,js,templ,go}",
  ],
  darkMode: "class",
  theme: {
    extend: {
      fontFamily: {
        mono: ["Courier Prime", "monospace"],
      },
    },
  },
  plugins: [require("@tailwindcss/forms")],
  corePlugins: {
    preflight: true,
  },
};
