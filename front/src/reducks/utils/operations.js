import { fetchWrapper } from '../../utils/http';
import { showMessageAction } from '../messages/actions';
import {
  fetchPrefMasterAction,
} from './actions';

export const fetchPrefMaster = () => {
  return async (dispatch) => {
    fetchWrapper(
      {
        type: 'GET',
        url: '/master/pref',
      },
      dispatch,
    )
      .then((json) => {
        if (!json) return;
        if (json.status === 200) {
          const prefMaster = json.result.prefMaster;
          if (Array.isArray(prefMaster)) {
            dispatch(
              fetchPrefMasterAction(prefMaster)
            );
          }
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
      });
  };
};
