/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html",
  "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        "color-primary": "#d52a47",
        "color-secondary": "#df5134",
        "color-primary-bg": "#FCFAF7",
        "color-text-header": "#292828",
        "color-text-body": "#373434",
        "color-text-details": "#b3b2b2",
      }
    },
    fontFamily: {
      Roboto: ["Roboto", "sans-serif"],
    },
    container: {
      padding: "2rem",
      center: true,
    }
  },
  plugins: [],
}
