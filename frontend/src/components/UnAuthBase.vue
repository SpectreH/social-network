<template>
  <section class="sign-in-page">
    <div id="container-inside">
      <div id="circle-small"></div>
      <div id="circle-medium"></div>
      <div id="circle-large"></div>
      <div id="circle-xlarge"></div>
      <div id="circle-xxlarge"></div>
    </div>
    <div class="container p-0">
      <div class="row no-gutters justify-content-center align-items-center h-100">
        <div class="col-md-8 bg-white pt-5 pt-5 pb-5">
          <router-view></router-view>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: "UnAuthBase",
  data: () => ({}),
};
</script>
