<script lang="ts">
  import {
    faCirclePlus,
    faDownload,
    faKey,
    faPlusCircle,
    faTrash,
    faUserPlus,
    faUserShield,
    faXmark,
  } from "@fortawesome/free-solid-svg-icons";
  import "../../app.scss";
  import UserManagementBox from "../users/UserManagementBox.svelte";
  import Fa from "svelte-fa";
  import { whoami } from "../../script/whoami";
  import { Button, Modal, Spinner } from "flowbite-svelte";
  import { searchIcons } from "../../script/icongen";
  import {
    listImages,
    type ListedImage,
    type ListedImagesGrouped,
  } from "../../script/images-networks-volumes";
  import ImageBox from "./ImageBox.svelte";

  let { pushAlert, updatePage } = $props();

  function groupImagesByTag(ungrouped: ListedImage[]): ListedImagesGrouped {
    const grouped: ListedImagesGrouped = {};

    for (const image of ungrouped) {
      // Create entry if non existent
      if (!grouped[image.Repository])
        grouped[image.Repository] = {
          tags: [],
          images: [],
        };

      // Add image to grouped
      grouped[image.Repository].images.push(image);

      // Add tag
      grouped[image.Repository].tags.push(image.Tag);
    }

    return grouped;
  }

  let showPullPopup = $state(false);
  let refetch = $state(Math.random());

  function refetchImages() {
    refetch = new Date().getTime() ** Math.random();
  }
</script>

<main
  class="flex justify-start flex-col w-full h-full items-start pt-4 pb-4 pl-8 pr-8"
>
  <!-- Heading + Create User Button -->
  <h1 class="text-3xl mb-8 heading">
    <span class="l">Images</span>

    {#await whoami() then userInfo}
      {#if userInfo.admin}
        <!-- svelte-ignore a11y_click_events_have_key_events -->
        <!-- svelte-ignore a11y_no_static_element_interactions -->
        <span class="r" onclick={() => (showPullPopup = true)}>
          <Fa icon={faPlusCircle} />
        </span>
      {/if}
    {/await}
  </h1>

  <!-- Images List -->
  {#key refetch}
    {#await listImages()}
      <div class="spinner-center text-center">
        <Spinner size={20} />
        <h2 class="mt-5">Loading images, please wait...</h2>
      </div>
    {:then images}
      {#each Object.entries(groupImagesByTag(images)).sort( (a, b) => ("" + a[0]).localeCompare(b[0]) ) as [repository, imagesWithSameTag]}
        <ImageBox
          {pushAlert}
          {updatePage}
          {imagesWithSameTag}
          {refetchImages}
        />
      {/each}
    {/await}
  {/key}
</main>

<!-- Local styling -->
<style lang="scss">
  .heading {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    width: 100%;

    .r {
      color: #102e4a;
      &:hover {
        cursor: pointer;
        color: #0d7e2b;
      }
    }
  }

  .spinner-center {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    width: 100%;
    height: 100%;
  }
</style>
