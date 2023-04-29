import * as Actions from './actions';
import initialState from '../store/initialState';

export const UtilsReducer = (state = initialState.utils, action) => {
  switch (action.type) {
    case Actions.WINDOW_RESIZE:
      return {
        ...state,
        windowSize: {...action.payload},
      };
    case Actions.FETCH_PREF_MASTER:
      return {
        ...state,
        prefMaster: [...action.payload],
      };
    case Actions.CUSTOM_VALID_ERR:
      return {
        ...state,
        validation: {
          error: true,
          errorResult: {...action.payload}
        },
      };
    case Actions.CUSTOM_VALID_ERR_RESET:
      return {
        ...state,
        validation: {
          error: false,
          errorResult: null,
        },
      };
    default:
      return state;
  }
};
