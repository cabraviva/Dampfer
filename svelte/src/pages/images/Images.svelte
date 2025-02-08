<script lang="ts">
  import {
    faCirclePlus,
    faKey,
    faTrash,
    faUserPlus,
    faUserShield,
    faXmark,
  } from "@fortawesome/free-solid-svg-icons";
  import "../../app.scss";
  import UserManagementBox from "../users/UserManagementBox.svelte";
  import Fa from "svelte-fa";
  import { whoami } from "../../script/whoami";
  import {
    createUser,
    deleteUser,
    listUsers,
    setPasswordForUser,
    setPermissionForUser,
    type UserList,
  } from "../../script/user-api";
  import SmallPopup from "../../popups/SmallPopup.svelte";
  import { Button, Modal } from "flowbite-svelte";

  let { pushAlert, updatePage } = $props();

  let users: UserList = $state([]);

  function refetchUsers() {
    listUsers().then((userVal) => {
      users = userVal;
    });
  }

  refetchUsers();

  let showPullPopup = $state(false);
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
          <Fa icon={faCirclePlus} />
        </span>
      {/if}
    {/await}
  </h1>

  <!-- Images List -->
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
</style>
