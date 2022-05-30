import store from '@/store'

store.subscribe((mutation) => {
  switch (mutation.type) {
    case 'auth/SET_TOKEN':
      if (mutation.payload) {
        localStorage.setItem('sn_token', mutation.payload)
      } else {
        localStorage.setItem('sn_token', "")
      }
      break;
  }
})
