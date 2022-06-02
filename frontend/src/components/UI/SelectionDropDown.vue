<template>
  <div class="btn-group">
    <button type="button" class="btn btn-primary dropdown-toggle" data-bs-toggle="dropdown" :style="additionalButton ? 'border-top-left-radius: 0; border-bottom-left-radius: 0' : ''">{{selectionAttr.label}}</button>
    <div class="dropdown-menu" onclick="event.stopPropagation()">
      <div v-if="selectionAttr.elements.length !== 0">
        <div v-for="(element, index) of selectionAttr.elements" :key="index">
          <div class="mx-2 dropdown-element">
            <input type="checkbox" :id="element.id" v-model="element.selected">
            <label :for="element.id">{{element.firstName}} {{element.lastName}}</label>
          </div>
          <hr class="m-0">
        </div>
      </div>
      <div v-else>
        <p class="m-0 text-center">You don't have followers :(</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SelectionDropDown',
  props: {
    selectionAttr: {type: Object, default: () => { 
      return {
        label: "Selection",
        elements: [],
      } 
    }},
    additionalButton: {type: Boolean, default: false}
  }
}
</script>

<style scoped>
.dropdown-element {
  display: flex;
  align-items: center;
  gap: 20px;
  background-position-y: center;
}

.dropdown-menu {
  max-height: 200px;
  max-width: 250px;
  overflow: overlay;  
}
</style>