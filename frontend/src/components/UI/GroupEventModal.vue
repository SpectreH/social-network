<template>
  <div class="modal fade" :id="modalId" tabindex="-1" aria-labelledby="modalLabel" aria-hidden="true" style="overflow: inherit">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="modalLabel">Event name</h5>
          <button
            type="button"
            class="btn btn-secondary"
            data-bs-dismiss="modal"
          >
            <i class="ri-close-line"></i>
          </button>
        </div>

        <div class="modal-body">
          <div class="d-flex justify-content-center">
            <div class="profile-img-edit">
              <img class="profile-pic" src="https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/user/11.png" alt="profile-pic">
            </div>
          </div>
          <hr>

          <div class="d-flex gap-2 mb-2">
            <h6 class="">Date:</h6>
            <h6 class="">28.03.02 2002</h6>
          </div>
          <hr>

          <div class="d-grid align-items-center gap-2 mb-2">
            <h6 class="">Description:</h6>
            <h6 class="">Is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</h6>
          </div>

          <hr>          
          <div class="d-flex justify-content-between">
            <div class="d-flex gap-2">
              <h6 class="">Will participate:</h6>
              <h6 class="">100</h6>
            </div>
            <div class="d-flex gap-2">
              <h6 class="">Refused:</h6>
              <h6 class="">100</h6>
            </div>
          </div>
        </div>
        <div class="modal-footer justify-content-center">
          <button type="button" class="col-4 btn btn-primary">Will participate</button>
          <button type="button" class="col-4 btn btn-primary">Refuse</button>
          <button type="button" class="col-4 btn btn-danger">Remove</button>
        </div>  
      </div>
    </div>
  </div>  
</template>

<script>
export default {
  name: "GroupEventModal",
  props: {
    modalId: {type: String, default: ""}
  },
  data: () => ({    
    newEvent: {
      title: "",
      description: "",
      time: null,
    },
  })    
}
</script>

<style scoped>
textarea {
  resize: none;
  overflow: hidden;
  min-height: 30px;
  max-height: 120px;
}
</style>