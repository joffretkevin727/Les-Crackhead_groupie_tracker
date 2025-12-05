import { createWeb3Modal, defaultConfig } from 'https://esm.sh/@web3modal/ethers@3.5.0'

const projectId = '866da356e8fbff41de2ba66db182553e'

const mainnet = {
  chainId: 1,
  name: 'Ethereum',
  currency: 'ETH',
  explorerUrl: 'https://etherscan.io',
  rpcUrl: 'https://cloudflare-eth.com'
}

const metadata = {
  name: 'Lizard Crypto',
  description: 'Crypto currencies exchange',
  url: 'https://monsite.com',
  icons: ['https://avatars.mywebsite.com/']
}

const ethersConfig = defaultConfig({
  metadata,
  enableEIP6963: true,
  enableInjected: true,
  enableCoinbase: true,
})

const modal = createWeb3Modal({
  ethersConfig,
  chains: [mainnet],
  projectId,
  enableAnalytics: true
})

const btn = document.getElementById("account-btn");

if (btn) {
  btn.onclick = () => {
    modal.open()
  }
} else {
  console.error("Erreur: Le bouton avec l'ID 'account-btn' est introuvable dans le HTML.");
}

modal.subscribeState(state => {
  const isConnected = state.selectedNetworkId !== undefined;
  console.log("État connexion:", isConnected);

  // if (isConnected && btn) {
  //   btn.innerText = "CONNECTÉ !";
  // }
})

// 1. Écouter les changements de compte (Connexion, Déconnexion, Changement de compte)
modal.subscribeProvider((state) => {
  if (state.address) {
    const userAddress = state.address;
    console.log("L'utilisateur est connecté avec :", userAddress);
    setBtnAddress(userAddress);

    fetch('/api/save-wallet', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ address: userAddress })
    });
  } else {
    console.log("Utilisateur déconnecté");
    setBtnAddress(null);
  }
});


try {
  // On demande au modal : "Est-ce qu'on a déjà une adresse en mémoire ?"
  const currentAddress = modal.getAddress();

  if (currentAddress) {
    console.log("✅ Déjà connecté au démarrage avec :", currentAddress);

    // On met à jour le bouton tout de suite
    const btn = document.getElementById("account-btn");
    if (btn) {
      btn.innerText = currentAddress.slice(0, 6) + "..." + currentAddress.slice(-4);
    }
  } else {
    console.log("⚪ Pas connecté au démarrage.");
  }
} catch (error) {
  // Parfois, si le provider n'est pas encore prêt, ça peut échouer silencieusement
  console.log("Attente de l'initialisation...");
}

function setBtnAddress(address) {
  if (!btn) return;

  if (!address) {
    btn.innerText = "MY ACCOUNT";
    return;
  }

  const short = address.slice(0, 6) + "..." + address.slice(-4);
  btn.innerText = short;
}
