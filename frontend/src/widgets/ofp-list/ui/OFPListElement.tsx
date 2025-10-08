import type { OFP } from "../../../entities/OFP";
import styles from "./OFPListElement.module.css";

export const OFPListElement = ({ ofp }: { ofp: OFP }) => {
  return (
    <div className={styles.ofpListElement}>
      <div>{ofp.icaoFrom}</div>
      <div>{ofp.icaoTo}</div>
      <div>{ofp.ETD}</div>
      <div>{ofp.ETA}</div>
      <div>{ofp.flightNumber}</div>
      <div>{ofp.dof}</div>
      <div>{ofp.regNumber}</div>
    </div>
  );
};
