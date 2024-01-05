<script>
  import FileSpace from "./components/FileSpace.svelte";

  let fileSpaceCount = 1;
</script>

<main class="main-container">
  
  <h2 class="header">knucks</h2>
  <div class="file-space">
    {#each [...Array(fileSpaceCount).keys()] as i}
      <FileSpace />
    {/each}
  </div>
  
  
</main>

<style>
  main {
    --header-height: 30px;
  }
  .header {
    padding: 0;
    margin: 0;
    height: var(--header-height);
  }
  .file-space {
    height: calc(100vh - var(--header-height));
  }
  .main-container {
    width: 100vw;
    height: 100%;
    min-height: 100%;
    box-sizing: border-box;
    overflow-x: hidden;
    overflow-y: hidden;
  }
</style>
