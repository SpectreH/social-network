<template>
  <div class="card-body text-center">
    <div>
      <img :src="group.image" alt="profile-img" class="rounded-circle img-fluid avatar-120">
    </div>
    <div class="group-info pt-3 pb-3">
      <h4><router-link :to="'/group/'+ group.id">{{group.title}}</router-link></h4>
    </div>
    <div class="group-details d-inline-block pb-3">
      <ul class="d-flex align-items-center justify-content-between list-inline m-0 p-0">
        <li class="pe-3 ps-3">
          <p class="mb-0">Posts</p>
          <h6>{{group.posts}}</h6>
        </li>
        <li class="pe-3 ps-3">
          <p class="mb-0">Members</p>
          <h6>{{group.members}}</h6>
        </li>
        <li class="pe-3 ps-3">
          <p class="mb-0">Type</p>
          <h6>{{group.type.charAt(0).toUpperCase() + group.type.slice(1)}}</h6>
        </li>
      </ul>
    </div>
    <router-link class="btn btn-primary d-block w-100" :to="'/group/'+ group.id">Explore</router-link>
  </div>
</template>

<script>
export default {
  name: "GroupOverview",
  props: {
    group: {type: Object, default: () => {
      return {}
    }
  }}
}
</script>

<style>

</style>