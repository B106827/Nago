import * as Actions from './actions';
import initialState from '../store/initialState';

export const MessageReducer = (state = initialState.message, action) => {
  switch (action.type) {
    case Actions.SHOW_MESSAGE:
      return {
        ...state,
        ...action.payload,
      };
    case Actions.HIDE_MESSAGE:
      return {
        ...state,
        ...action.payload,
      };
    default:
      return state;
  }
};
