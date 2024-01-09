<script>
    import WaveSurfer from "wavesurfer.js";
    import SamplePlayer from "./SamplePlayer.svelte";
    import { CurrentDir } from "../../wailsjs/go/main/App.js"
    import { ReturnDirItems } from "../../wailsjs/go/main/App.js"
    import { MoveDir } from "../../wailsjs/go/main/App.js"
    import { SetStartingDir } from "../../wailsjs/go/main/App.js"
    let currentDirectory = []
    let allDirectoryItems = []
    let shownItems = []
    let browserCols = 2;
    $: loaded_sample = "";
    let searchQuery = '';
    let searching = false;
    function init() {
      getCurrentDir();
    }
  
    init();
  
    function convertStringToArray(input_string) {
      try {
        const jsonArray = JSON.parse(input_string);
        return jsonArray
      } catch (e) {
        console.error('error parsing string:', e)
        return null
      }
    }
  
    function getCurrentDir() {
      CurrentDir().then((res) => (currentDirectory = res)).then(() => {
        ReturnDirItems(currentDirectory).then((res) => {
          allDirectoryItems = convertStringToArray(res)
          shownItems = allDirectoryItems;
        })
      })
    }
  
    function MoveCurrentDir(newDir) {
      MoveDir(newDir).then(() => getCurrentDir())
      resetSearch();
    }

    function handleDragEnd(e) {
        if (windowHeight == 100) {
            windowHeight = 100;
        } else {
            windowHeight = 100;
        }
    }
    let windowWidth = 100;
    let windowHeight = 100;

    function handleSearch() {
      searching = true;
      searchItems()
    }

    function resetSearch() {
      shownItems = allDirectoryItems;
      searchQuery = ''
      searching = false;
    }

    function searchItems() {
      if (!searchQuery) {
        resetSearch()
        return
      }
      
      const searchedItems = allDirectoryItems.filter((item) => {
        return item.name.includes(searchQuery)
      })

      shownItems = searchedItems ?? allDirectoryItems;
      searching = false;
    }

    async function getDir() {
      SetStartingDir().then(() => getCurrentDir());
    }

  
    let isDragged = false;
    let curDrag = -1;
  </script>
  <div 
    
    class="window" 
    style="--browser-columns: {browserCols}; --window-height: {windowHeight}%; --window-width: {windowWidth}%;"
    on:dragstart={
        () => { windowHeight = windowHeight }
    }
    on:dragend={handleDragEnd}
    >
    <div class="dir-container">
        {#each currentDirectory as dir}
        <button class="dir-part" on:click={() => MoveCurrentDir(dir)}>
            {dir}
        </button>
        {/each}
        <div style="position: absolute; width: 300px; right: 0px; top: 7px;">
          <button style="float: right;" on:click={getDir}>select directory</button>
          <button style="float: right;" on:click={() => {
              if (browserCols > 0) {
                browserCols--
              }
              
            }}>-</button>
          <button style="float: right;" on:click={() => {
              if (browserCols < 4) {
                browserCols++
              }
              
            }}>+</button>
          <input style="float: right;" id="search" bind:value={searchQuery} on:input={handleSearch} />
          
        </div>
    </div>
    <div class="container-box">
        
          <div class="item-container" >
            {#each shownItems as item, index}
              {#if item.is_directory}
                <button class="dir-folder" on:click={() => MoveCurrentDir(item.name)}>
                  {item.name}
                </button>
              {:else}
                  
                  {#if item.is_audio}
                    <button
                    draggable="true"
                    on:click={() => { 
                      loaded_sample = item.location 
                      console.log(loaded_sample)
                    }}
                      >
                      {item.name}
                    </button>
                    {:else}
                    <button
                    draggable="true"
                    
                      >
                      {item.name}
                    </button>
                  {/if}
              {/if}
            {/each}
          </div>
    </div>
    <div id="waveform">

    </div>
  </div>
    
  
  <style>
    .window {
        height: var(--window-height);
    }
    .dir-part {
      background-color: rgba(169, 169, 169, 0.932);
      font-size: 16px;
      border: none;
      border-radius: 0;
      cursor: pointer;
      text-align: left;
      /* max-height: 20px; */
    }
    .is-dragged {
      background-color: black;
      border: dashed 2px gray;
    }
    .dir-part:hover {
      opacity: 0.8;
      transition: 0.05s;
    }
  
    .dir-container {
    background-color: rgba(128, 255, 0, 0.342);
      display: flex;
      width: 100%;
    }
  
    .dir-item, .dir-folder {
      background-color: white;
      text-align: left;
      font-size: 18px;
      
      border-radius: 0;
      cursor: pointer;
      transition: 0.1s;
    }
  
    .dir-item:hover, .dir-folder:hover {
      
      opacity: 0.8;
      transition: 0.06s;
    }

    .dir-item:hover {
      background-color: rgba(184, 119, 119, 0.151);
    }

    .dir-folder:hover {
      background-color: rgba(119, 133, 133, 0.151);
    }
  
    .dir-folder {
      color: brown;
      
    }
    .dir-item {
      color: cadetblue;
      
    }
    .item-container {
      display: grid;
      padding: 2px;
      grid-template-columns: repeat(var(--browser-columns), 1fr);
    }

    .container-box {
        height: 100%;
        border: solid 2px rgba(122, 122, 122, 1);
        overflow: auto;
    }
  
  </style>
  