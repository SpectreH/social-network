<template>
  <a
    href="#"
    class="dropdown-toggle d-grid"
    data-bs-toggle="dropdown"
    aria-haspopup="true"
    aria-expanded="false"
    :id="menuAttr.id"
  >
    <i :class="menuAttr.icon"></i>
    <div v-if="menuAttr.elements.length !== 0" class="number-circle danger" style="font-size: 11px">{{ menuAttr.elements.length }}</div>
  </a>

  <div :class="menuAttr.id === 'group-drop' ? 'sub-drop-large' : _ " class="sub-drop dropdown-menu" :aria-labelledby="menuAttr.id">
    <div class="card shadow-none m-0">
      <div
        class="
          card-header
          d-flex
          justify-content-between
          bg-primary
        "
      >
        <div class="header-title bg-primary">
          <h5 class="mb-0 text-white">{{ menuAttr.title }}</h5>
        </div>
        <small class="badge bg-light text-dark">{{ menuAttr.elements.length }}</small>
      </div>
      <div class="card-body p-0">

        <div v-if="menuAttr.id === 'notification-drop'">
          <a v-for="element in menuAttr.elements" :key="element.id" href="#" class="iq-sub-card">
            <div class="d-flex align-items-center">
              <div class="">
                <img :src="element.avatar" class="avatar-40 rounded" alt="" />
              </div>
              <div class="ms-3 w-100">
                <h6 class="mb-0">{{ element.author }}</h6>
                <div
                  class="
                    d-flex
                    justify-content-between
                    align-items-center
                  "
                >
                  <p class="mb-0">{{ element.sub }}</p>
                  <small class="float-right font-size-12">{{ element.time }}</small>
                </div>
              </div>
            </div>
          </a>
          <div v-if="menuAttr.elements.length === 0">
            <p class="text-center mb-0">You don't have any notification</p>     
          </div>
        </div>

        <div v-else-if="menuAttr.id === 'group-drop'">
          <div v-for="element in menuAttr.elements" :key="element.id" class="iq-friend-request">
            <div
              class="
                iq-sub-card iq-sub-card-big
                d-flex
                align-items-center
                justify-content-between
              "
            >
              <div class="col-6 d-flex align-items-center">
                <img :src="element.avatar" class="avatar-40 rounded" alt="" />
                <div class="ms-3">
                  <h6 class="mb-0">{{ element.author }}</h6>
                  <p class="mb-0">{{ element.sub }}</p>
                </div>
              </div>
              <div class="col-5 d-flex align-items-center">
                <a
                  href="javascript:void();"
                  class="me-3 btn btn-primary rounded"
                  >Confirm</a
                >
                <a
                  href="javascript:void();"
                  class="me-3 btn btn-secondary rounded"
                  >Decline</a
                >
              </div>
            </div>
          </div>
          <div v-if="menuAttr.elements.length === 0">
            <p class="text-center mb-0">You don't have any request</p>     
          </div>
        </div>

      </div>
    </div> 
  </div>
</template>

<script>
export default {
  name: "NavDropDown",
  props: {
    menuAttr: {type: Object, default: () => { 
      return {
        id: "",
        icon: "",
        title: "",
        elements: [],
      } 
    }}
  }
}
</script>

<style>
.number-circle {
  font: 32px Arial, sans-serif;
  position: absolute;

  width: 2em;
  height: 2em;
  box-sizing: initial;
  
  background-color: red;
  border: 0.1em solid #666;
  color: white;
  text-align: center;
  border-radius: 50%;    

  line-height: 2em;
  box-sizing: content-box;
  
  margin-top: 7px;
  margin-left: 13px;
}
</style>