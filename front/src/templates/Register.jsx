import { useCallback, useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { push } from 'connected-react-router';
import { useParams } from 'react-router-dom';
import { showMessageAction } from '../reducks/messages/actions';
import { checkRegisterUrl, register } from '../reducks/users/operations';
import { getUserTmpEmail } from '../reducks/users/selectors';
import { TextInput, PrimaryButton } from '../components/UIkit';

const Register = () => {
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);

  // URL が有効かどうか確認する
  const { tmpId } = useParams();
  if (!tmpId) {
    dispatch(showMessageAction('error', '無効なURLです'));
    dispatch(push('/'));
  } else {
    useEffect(() => {
      dispatch(checkRegisterUrl(tmpId));
    }, []);
  }

  const [name, setName] = useState(''),
    [email, setEmail] = useState(''),
    [password, setPassword] = useState(''),
    [confirmPassword, setConfirmPassword] = useState('');

  const inputName = useCallback(
    (event) => {
      setName(event.target.value);
    },
    [setName]
  );

  // メールアドレスは仮登録アドレスを関ししておき、あればそれを入れる
  const userTmpEmail = getUserTmpEmail(selector);
  useEffect(() => {
    if (userTmpEmail) setEmail(userTmpEmail);
  }, [userTmpEmail]);

  const inputPassword = useCallback(
    (event) => {
      setPassword(event.target.value);
    },
    [setPassword]
  );

  const inputConfirmPassword = useCallback(
    (event) => {
      setConfirmPassword(event.target.value);
    },
    [setConfirmPassword]
  );

  return (
    <div className='c-section-container'>
      <h2 className='u-text__headline__k u-text-center'>新規登録</h2>
      <div className='module-spacer--medium' />
      <TextInput
        fullWidth={true}
        label={'お名前'}
        multiline={false}
        required={true}
        rows={1}
        value={name}
        type={'text'}
        onChange={inputName}
      />
      <TextInput
        fullWidth={true}
        label={'メールアドレス'}
        multiline={false}
        required={true}
        rows={1}
        value={email}
        type={'email'}
        variant={'filled'}
        InputProps={{
          readOnly: true,
        }}
      />
      <TextInput
        fullWidth={true}
        label={'パスワード'}
        multiline={false}
        required={true}
        rows={1}
        value={password}
        type={'password'}
        onChange={inputPassword}
      />
      <TextInput
        fullWidth={true}
        label={'パスワード(再確認)'}
        multiline={false}
        required={true}
        rows={1}
        value={confirmPassword}
        type={'password'}
        onChange={inputConfirmPassword}
      />
      <div className='module-spacer--medium' />
      <div className='center'>
        <PrimaryButton
          label={'アカウントを登録する'}
          onClick={() =>
            dispatch(register(tmpId, name, email, password, confirmPassword))
          }
          addStyle={{ fontWeight: 'bold' }}
        />
        <div className='module-spacer--medium' />
        <p className='cursor-pointer' onClick={() => dispatch(push('/login'))}>
          アカウントをお持ちの方はこちら
        </p>
      </div>
    </div>
  );
};

export default Register;
