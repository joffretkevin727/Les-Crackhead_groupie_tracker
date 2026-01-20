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
//on exporte la constante modal pour pouvoir gerer les deconnections dans profil
export const modal = createWeb3Modal({
  ethersConfig,
  chains: [mainnet],
  projectId,
  enableAnalytics: true
})

const btn = document.getElementById("account-btn");
const logoutBtn = document.getElementById('logout-btn');

if (btn) {
  btn.onclick = () => {
    if (btn.innerText.trim() === "CONNEXION") {
      modal.open();
    } else {
      window.location.href = "/profil";
    }

  }
} else {
  console.error("Erreur: Le bouton avec l'ID 'account-btn' est introuvable dans le HTML.");
}

// 2. GESTION DU BOUTON DE DÉCONNEXION (PAGE PROFIL UNIQUEMENT)
if (logoutBtn) {
    logoutBtn.onclick = async () => {
        await modal.disconnect();
        window.location.href = "/home"; 
    };
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
    const logoutBtn = document.getElementById('logout-btn');
    const walletDisplay = document.getElementById("wallet-address-display");

    if (state.address) {
        // État Connecté
        if (btn) btn.innerText = "MY ACCOUNT";
        if (walletDisplay) walletDisplay.innerText = state.address;
        if (logoutBtn) logoutBtn.style.display = "block"; // Affiche Logout

        fetch('/api/save-wallet', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ address: state.address })
        });
    } else {
        // État Déconnecté
        if (btn) btn.innerText = "CONNEXION";
        if (walletDisplay) walletDisplay.innerText = "";
        if (logoutBtn) logoutBtn.style.display = "none"; // Cache Logout
    }
});


try {
    const currentAddress = modal.getAddress();
    
    if (currentAddress) {
        // 1. Mettre à jour le bouton de navigation
        if (btn) btn.innerText = "MY ACCOUNT";
        
        // 2. Gérer les éléments spécifiques à la page profil
        const logoutBtn = document.getElementById('logout-btn');
        const walletDisplay = document.getElementById("wallet-address-display");
        
        if (logoutBtn) logoutBtn.style.display = "block"; // Affiche Logout
        if (walletDisplay) walletDisplay.innerText = currentAddress; // Affiche l'adresse
    }
} catch (error) {
    console.log("Initialisation du provider en cours...");
}

function setBtnAddress(address) {
  if (!btn) return;
  btn.innerText = address ? "MY ACCOUNT" : "CONNEXION";
}
