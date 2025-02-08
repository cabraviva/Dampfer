<script lang="ts">
  import Fa from "svelte-fa";
  import "../../app.scss";
  import "./image-box.scss";
  import type { AlertType } from "../../types";
  import type { ListedImage } from "../../script/images-networks-volumes";
  import { onMount } from "svelte";
  import { getAverageBackgroundColor, searchIcons } from "../../script/icongen";
  import knorry from "knorry";
  import { getCredentials } from "../../script/login";

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

<!-- TODO:
        - Pull images
        - Authentication when pulling
        - Editing image: Delete, change image via A) Search B) Upload
        - Better image info on list page
        - Detailed image info when clicking on image (using image inspect)
-->

<div class="image-box">
  <div class="icon" bind:this={iconContainer}>
    {#await fetch( `/api/icongen/get-icon?id=${encodeURIComponent(image.ID)}`, { headers: { Authorization: `Bearer ${getCredentials()}` } } ) then imgRes}
      {#await imgRes.blob() then blob}
        <img src={URL.createObjectURL(blob)} alt="" />
      {/await}
    {/await}
  </div>
  <div class="info">
    {image.Repository}:{image.Tag}
  </div>
  <div class="buttons"></div>
</div>
