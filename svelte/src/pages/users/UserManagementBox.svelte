<script lang="ts">
  import Fa from "svelte-fa";
  import "../../app.scss";
  import "./user-man.scss";
  import {
    faCrown,
    faKey,
    faShieldHalved,
    faTrash,
    faUser,
    faUserTie,
  } from "@fortawesome/free-solid-svg-icons";

  let {
    pushAlert,
    updatePage,
    username,
    permission,
    isme = false,
    onChangePermission,
    onChangePassword,
    onDelete,
  } = $props();

  let permissionText = $state("Insight");

  if (permission === "admin") permissionText = "Admin";
  if (permission === "system-admin") permissionText = "System Admin";
</script>

<div class="user-box">
  <div class="user-left">
    {#if permission === "system-admin"}
      <span class="user-icon system-admin"><Fa icon={faCrown} /></span>
    {:else if permission === "admin"}
      <span class="user-icon admin"><Fa icon={faUserTie} /></span>
    {:else}
      <span class="user-icon insight"><Fa icon={faUser} /></span>
    {/if}

    <span class="user-name">{username}</span>

    <span class="user-perm">{permissionText}</span>
  </div>

  <div class="user-right">
    <button
      class="user-btn-chperm user-btn"
      title="Change permission"
      aria-label="Change permission"
      onclick={onChangePermission}
    >
      <Fa icon={faShieldHalved} />
    </button>

    <button
      class="user-btn-chpwd user-btn"
      title="Change password"
      aria-label="Change password"
      onclick={onChangePassword}
    >
      <Fa icon={faKey} />
    </button>

    {#if !isme}
      <button
        class="user-btn-del user-btn"
        title="Delete user"
        aria-label="Delete user"
        onclick={onDelete}
      >
        <Fa icon={faTrash} />
      </button>
    {/if}
  </div>
</div>
