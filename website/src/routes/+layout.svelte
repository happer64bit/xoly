<script>
  import Header from "./Header.svelte";
  import "./../app.css";
  import "@fontsource/poppins/400.css";
  import "@fontsource/poppins/500.css";
  import "@fontsource/poppins/600.css";
  import "@fontsource/poppins/700.css";
</script>

<Header />
<slot />

<style>
  :global(body) {
    font-family: "Poppins", sans-serif;
  }
</style>
