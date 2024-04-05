<script lang="ts">
  import "@fontsource/fira-mono/500.css";
  
  export let data;
</script>

<div>
  <div class="h-[80vh] flex items-center justify-center">
    <div>
      <div class="text-center my-10">
        <h2 class="text-4xl font-bold">Preview Shortened URL</h2>
      </div>
      <div class="w-[80dvh] grid grid-cols-6 gap-4">
        <label for="" class="col-span-2">Shortened URL</label>
        <input
          type="text"
          class="bg-gray-50 px-4 py-2 h-11 w-full outline-none border-2 focus:border-blue-500 rounded-lg transform duration-200 font-fira col-span-4"
          placeholder="https://..."
          value={data.slug}
          disabled
        />
        <label for="" class="col-span-2">Link</label>
        <input
          type="text"
          class="bg-gray-50 px-4 py-2 h-11 w-full outline-none border-2 focus:border-blue-500 rounded-lg transform duration-200 font-fira col-span-4"
          placeholder="https://..."
          value={data.url}
          disabled
        />
      </div>
    </div>
  </div>
</div>

<style>
  .font-fira {
    font-family: "Fira Code", monospace;
  }
</style>
