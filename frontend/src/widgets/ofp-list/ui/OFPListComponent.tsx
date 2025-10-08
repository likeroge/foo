import { useEffect, useState } from "react";
import styles from "./OFPListComponent.module.css";
import type { OFP } from "../../../entities/OFP";
import { getAllOFP } from "../api/get-all-ofp";
import { OFPListElement } from "./OFPListElement";

export const OFPListComponent = () => {
  const [ofpList, setOFPList] = useState<OFP[]>([]);
  useEffect(() => {
    (async () => {
      try {
        const ofpList = await getAllOFP();
        setOFPList(ofpList);
      } catch (error) {
        console.error("ERROR: ", (error as Error).message);
      }
    })();
  }, []);
  return (
    <div className={styles.ofpList}>
      {ofpList.map((ofp) => (
        <OFPListElement key={ofp.id} ofp={ofp} />
      ))}
    </div>
  );
};
