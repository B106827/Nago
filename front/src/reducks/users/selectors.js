import { createSelector } from 'reselect';

const usersSelector = (state) => state.users;

export const getIsLogedIn = createSelector(
  [usersSelector],
  (state) => state.isLogedIn
);

export const getOrdersHistory = createSelector(
  [usersSelector],
  (state) => state.orders
);

export const getProductsInCart = createSelector(
  [usersSelector],
  (state) => state.cart
);

export const getUserId = createSelector([usersSelector], (state) => state.id);

export const getUserName = createSelector(
  [usersSelector],
  (state) => state.name
);

export const getUserTmpEmail = createSelector(
  [usersSelector],
  (state) => state.tmpEmail
);
