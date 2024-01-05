<script>
    import { CurrentDir } from "../../wailsjs/go/main/App.js"
    import { ReturnDirItems } from "../../wailsjs/go/main/App.js"
    import { MoveDir } from "../../wailsjs/go/main/App.js"
    let currentDirectory = []
    let directoryItems = []
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
          directoryItems = convertStringToArray(res)
        })
      })
    }
  
    function MoveCurrentDir(newDir) {
      MoveDir(newDir).then(() => getCurrentDir())
      
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
  </script>
  <div 
    draggable="true" 
    class="window" 
    style="--window-height: {windowHeight}%; --window-width: {windowWidth}%;"
    on:dragstart={
        () => { windowHeight = windowHeight / 2 }
    }
    on:dragend={handleDragEnd}
    >
    <div class="dir-container">
        {#each currentDirectory as dir}
        <button class="dir-part" on:click={() => MoveCurrentDir(dir)}>
            {dir}
        </button>
        {/each}
    </div>
    <div class="container-box">
        
          <div class="item-container">
            {#each directoryItems as item}
            {#if item.is_directory}
              <button class="dir-folder" on:click={() => MoveCurrentDir(item.name)}>
                {item.name}
              </button>
            {:else}
              <button draggable="true" class="dir-item" on:click={() => MoveCurrentDir(item.name)}>
                {item.name}
              </button>
            {/if}
            {/each}
          </div>
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
      border: none;
      border-radius: 0;
      cursor: pointer;
      transition: 0.1s;
    }
  
    .dir-item:hover, .dir-folder:hover {
      opacity: 0.8;
      transition: 0.06s;
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
      grid-template-columns: 1fr 1fr;
      
    }

    .container-box {
        height: 100%;
        border: solid 2px rgba(122, 122, 122, 1);
        overflow: auto;
    }
  
  </style>
  