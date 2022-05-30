export default function guest({ next, store }) {
  if (store.getters['auth/authenticated']) {
    return next({
      name: 'home'
    })
  }

  return next();
}