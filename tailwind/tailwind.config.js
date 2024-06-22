/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["../views/**/*.{templ,go}"],
  theme: {
    extend: {},
  },
  plugins: [
    require("daisyui")
  ],
  daisyui: {
    themes: ["dark"]
  }
}
