import { useDispatch } from 'react-redux';
import { PrimaryButton } from '../components/UIkit';
import { push } from 'connected-react-router';

const NotFound = () => {
  const dispatch = useDispatch();

  return (
    <div className='c-section-container'>
      <h2 className='u-text__headline__k u-text-center'>404 Not Found</h2>
      <p className='u-text-center'>ページが見つかりません</p>
      <div className='module-spacer--medium' />
      <div className='center'>
        <PrimaryButton
          label={'トップに戻る'}
          onClick={() => dispatch(push('/'))}
        />
      </div>
    </div>
  );
};

export default NotFound;
