<script lang="ts">
  import Fa from "svelte-fa";
  import "../../app.scss";
  import "./image-box.scss";
  import type { AlertType } from "../../types";
  import type { ListedImage } from "../../script/images-networks-volumes";
  import { onMount } from "svelte";
  import { getAverageBackgroundColor, searchIcons } from "../../script/icongen";

  let { pushAlert, updatePage, image } = $props() as {
    pushAlert: (alert: AlertType) => void;
    updatePage: (name: string) => void;
    image: ListedImage;
  };

  let iconContainer: HTMLDivElement | null = $state(null);

  onMount(async () => {
    async function calcBg() {
      try {
        if (iconContainer !== null) {
          const avgBgColor = await getAverageBackgroundColor(
            iconContainer.querySelector("img") || document.createElement("img")
          );

          iconContainer.style.backgroundColor = avgBgColor;
        } else {
          setTimeout(calcBg, 50);
        }
      } catch {
        setTimeout(calcBg, 50);
      }
    }

    calcBg();
  });
</script>

<div class="image-box">
  <div class="icon" bind:this={iconContainer}>
    {#await searchIcons(image.Repository) then searchResults}
      <img src={searchResults[0]} alt="{image.Repository}:{image.Tag}" />
    {/await}
  </div>
  <div class="info">
    {image.Repository}:{image.Tag}
  </div>
  <div class="buttons"></div>
</div>
