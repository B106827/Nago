export const WINDOW_RESIZE = 'WINDOW_RESIZE';
export const windowResizeAction = (windowSize) => {
  return {
    type: WINDOW_RESIZE,
    payload: {
      width: windowSize.width,
      height: windowSize.height,
    },
  };
};

export const FETCH_PREF_MASTER = 'FETCH_PREF_MASTER';
export const fetchPrefMasterAction = (prefMaster) => {
  return {
    type: FETCH_PREF_MASTER,
    payload: prefMaster,
  };
};
