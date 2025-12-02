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
  
  if (isConnected && btn) {
    btn.innerText = "CONNECTÉ !";
  }
})