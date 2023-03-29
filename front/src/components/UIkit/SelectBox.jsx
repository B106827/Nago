import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import FormHelperText from '@material-ui/core/FormHelperText';
import Select from '@material-ui/core/Select';
import { makeStyles } from '@material-ui/styles';

const SelectBox = (props) => {
  const classes = useStyles();

  return (
    <FormControl required={props.required} className={classes.formControl} error={props.error} disabled={props.error}>
      <InputLabel className={classes.selectBoxLabel}>{props.label}</InputLabel>
      <Select
        className={classes.selectBox}
        value={props.value}
        onChange={(event) => props.select(event.target.value)}
      >
        {props.options && props.options.map((option) => (
          <MenuItem key={option.id} value={option.id}>
            {option.name}
          </MenuItem>
        ))}
      </Select>
      {props.error && (
        <FormHelperText className={classes.selectBoxLabel}>{props.errorMsg}</FormHelperText>
      )}
    </FormControl>
  );
};

const useStyles = makeStyles({
  formControl: {
    marginBottom: 16,
    minWidth: 128,
    width: '100%',
  },
  selectBox: {
    fontSize: 20
  },
  selectBoxLabel: {
    fontSize: 18
  }
});

export default SelectBox;
