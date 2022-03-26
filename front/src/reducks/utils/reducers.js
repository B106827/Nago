import * as Actions from './actions';
import initialState from '../store/initialState';

export const UtilsReducer = (state = initialState.utils, action) => {
  switch (action.type) {
    case Actions.WINDOW_RESIZE:
      return {
        ...state,
        windowSize: { ...action.payload },
      };
    default:
      return state;
  }
};
