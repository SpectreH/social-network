
<template>
  <div class="d-flex">
    <div class="user-img" style="min-width: 35px">
      <router-link to="/user/1">
        <img
          src="https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/user/02.jpg"
          alt="userimg"
          class="avatar-35 rounded-circle img-fluid"
          width="35px"
          height="35px"
        />
      </router-link>  
    </div>
    <div class="comment-data-block ms-3">
      <router-link to="/user/1">
        <h6 class="d-inline-flex">Monty Carlo</h6>
      </router-link>
      <p class="mb-0">Lorem ipsum dolor sit ametLorem ipsum dolor sit ametLorem ipsum dolor sit ametLorem ipsum dolor sit ametLorem ipsum dolor sit ametLorem ipsum dolor sit ametLorem ipsum dolor sit ametLorem ipsum dolor sit ametLorem ipsum dolor sit ametLorem ipsum dolor sit ametLorem ipsum dolor sit ametLorem ipsum dolor sit ametLorem ipsum dolor sit ametLorem ipsum dolor sit ametLorem ipsum dolor sit ametLorem ipsum dolor sit amet</p>
      <div
        class="d-flex flex-wrap align-items-center comment-activity"
      >
        <span class="text-primary"> 5 min </span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "PostComment"
}
</script>