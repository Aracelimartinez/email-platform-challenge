/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue, js, ts, jsx, tsx"],
  theme: {
    extend: {
      colors: {
        "primary-color": "#d52a47",
        "secondary-color": "#df5134",
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
