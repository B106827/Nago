import { createSelector } from 'reselect';

const utilsSelector = (state) => state.utils;

export const getWindowSize = createSelector(
  [utilsSelector],
  (state) => state.windowSize
);

export const getPrefMaster = createSelector(
  [utilsSelector],
  (state) => state.prefMaster
);

export const getValidErrFlg = createSelector(
  [utilsSelector],
  (state) => state.validation.error
);

export const getValidErrResult = createSelector(
  [utilsSelector],
  (state) => state.validation.errorResult
);
