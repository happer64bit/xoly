/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      container: {
        center: true,
        screens: {
          "2xl": "1300px"
        }
      }
    }
  },
  plugins: []
};
