/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.{html,js,templ,go}",
    "./templates/common/**/*.{html,js,templ,go}",
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
