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
