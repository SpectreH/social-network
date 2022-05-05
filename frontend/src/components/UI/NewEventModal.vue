<template>
  <div class="modal fade" :id="modalId" tabindex="-1" aria-labelledby="modalLabel" aria-hidden="true" style="overflow: inherit">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="modalLabel">Create Event</h5>
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
              <div class="p-image">
                <i class="ri-pencil-line upload-button text-white"></i>
                <input class="file-upload" type="file" accept="image/*">
              </div>
            </div>
          </div>
          <hr>
          
          <div class="d-flex align-items-center justify-content-between mb-2">
            <h6 class="">Title:</h6>
            <input type="text" class="form-control w-75" v-model="newEvent.title" placeholder="Event title">
          </div>
          <hr>

          <div class="d-flex align-items-center justify-content-between mb-2">
            <h6 class="">Date:</h6>
            <input type="date" class="form-control w-75" v-model="newEvent.time" placeholder="Event date">
          </div>
          <hr>

          <div class="d-flex align-items-center justify-content-between mb-2">
            <h6 class="">Description:</h6>
            <textarea id="descriptionTextArea" type="text" class="form-control w-75" v-model="newEvent.description" placeholder="Event description" style="resize: none; height: 60px"> </textarea>
          </div>
        </div>
        <div class="modal-footer justify-content-center">
          <button type="button" class="col-4 btn btn-primary">Create</button>
        </div>
      </div>
    </div>
  </div>  
</template>

<script>
export default {
  name: "NewEventModal",
  props: {
    modalId: {type: String, default: ""}
  },
  data: () => ({    
    newEvent: {
      title: "",
      description: "",
      time: null,
    },
  }),
  mounted() {
    const textarea = document.getElementById("descriptionTextArea");
    const limit = 120; // height limit

    textarea.oninput = function() {
      textarea.style.height = "";
      textarea.style.height = Math.min(textarea.scrollHeight, limit) + "px";
    };
  }    
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