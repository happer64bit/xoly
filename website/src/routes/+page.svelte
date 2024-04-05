<script>
  import { enhance } from "$app/forms";
  import Footer from "./Footer.svelte";
  import "@fontsource/fira-mono/500.css";
</script>

<div>
  <div class="h-[80vh] flex items-center justify-center">
    <div class="w-[90dvh]">
      <div class="text-center">
        <h1 class="text-3xl font-bold">Xoly: URL Shortener</h1>
      </div>
      <form class="mt-6 flex items-center" method="POST" action="?/create" use:enhance>
        <input
          type="text"
          name="url"
          class="bg-gray-50/50 px-4 py-2 h-11 w-full outline-none border-2 border-[#555] focus:border-blue-500 rounded-lg transform duration-200 font-fira placeholder:text-[#5e5e5e]"
          placeholder="https://..."
        />
      </form>
    </div>
  </div>
  <Footer />
</div>

<style>
  .font-fira {
    font-family: "Fira Code", monospace;
  }
</style>
